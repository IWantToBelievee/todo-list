package repository

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"sync"

	models "myproject/internals/models"
)

var encriptionKey = []byte("e047ce2fb1551a80206351b2274849a2")

func EncryptJSON(path string, data *models.TodosData) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		//log.Printf("Error marshaling JSON: %v", err)
		return err
	}
	//log.Printf("Marshaled JSON: %s", string(jsonData))

	block, err := aes.NewCipher(encriptionKey)
	if err != nil {
		return err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCTR(block, iv)
	encrypted := make([]byte, len(jsonData))
	stream.XORKeyStream(encrypted, jsonData)

	return os.WriteFile(path, append(iv, encrypted...), 0755)
}

func DecryptJSON(path string) (*models.TodosData, error) {
	encryptedData, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			// File not exist
			return CreateDefaultJSON(path)
		}
		return nil, err
	}

	if len(encryptedData) < aes.BlockSize {
		// Invalid encrypted file, creating new one
		return CreateDefaultJSON(path)
	}

	iv := encryptedData[:aes.BlockSize]
	encrypted := encryptedData[aes.BlockSize:]

	block, err := aes.NewCipher(encriptionKey)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, iv)
	decrypted := make([]byte, len(encrypted))
	stream.XORKeyStream(decrypted, encrypted)

	var data models.TodosData
	if err := json.Unmarshal(decrypted, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func CreateDefaultJSON(path string) (*models.TodosData, error) {
	defaultData := &models.TodosData{
		Todos: []*models.ToDo{},
		Order: []*string{},
	}

	_, err := os.Stat(filepath.Dir(path))
	if os.IsNotExist(err) {
		os.Mkdir(filepath.Dir(path), 0644)
	}

	if err := EncryptJSON(path, defaultData); err != nil {
		return nil, err
	}
	return defaultData, nil
}

func RemoveToDo(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		return slice // Проверка на валидность индекса
	}
	return append(slice[:i], slice[i+1:]...)
}

type JsonToDoRepository struct {
	path string
	mu   sync.RWMutex
}

func NewJsonToDoRepository(filename string) (*JsonToDoRepository, error) {
	confDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	return &JsonToDoRepository{path: filepath.Join(confDir, "todos-app", filename)}, nil
}

func (r *JsonToDoRepository) ReadTodosData() (*models.TodosData, error) {
	if data, err := DecryptJSON(r.path); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func (r *JsonToDoRepository) SaveTodosData(todosData *models.TodosData) error {
	if err := EncryptJSON(r.path, todosData); err != nil {
		return err
	}
	return nil
}

func (r *JsonToDoRepository) Create(todo *models.ToDo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := r.ReadTodosData()
	if err != nil {
		return err
	}

	data.AddTodo(todo)

	if err = r.SaveTodosData(data); err != nil {
		return err
	}

	return nil
}

func (r *JsonToDoRepository) GetAll() ([]*models.ToDo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	data, err := r.ReadTodosData()
	if err != nil {
		return nil, err
	}

	todos := data.Todos
	return todos, nil
}

func (r *JsonToDoRepository) GetByID(id *string) (*models.ToDo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	data, err := r.ReadTodosData()
	if err != nil {
		return nil, err
	}

	todo, _, err := data.GetByID(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *JsonToDoRepository) Update(req *models.UpdateRequest) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := r.ReadTodosData()
	if err != nil {
		return err
	}

	err = data.UpdateTodoByReq(req)
	if err != nil {
		return err
	}

	if err = r.SaveTodosData(data); err != nil {
		return err
	}

	return nil
}

func (r *JsonToDoRepository) Delete(id *string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := r.ReadTodosData()
	if err != nil {
		return err
	}

	err = data.RemoveTodo(id)
	if err != nil {
		return err
	}

	err = r.SaveTodosData(data)
	if err != nil {
		return err
	}

	return nil
}

func (r *JsonToDoRepository) SaveOrder(order []*string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := r.ReadTodosData()
	if err != nil {
		return err
	}

	data.Order = order

	if err = r.SaveTodosData(data); err != nil {
		return err
	}

	return nil
}

func (r *JsonToDoRepository) GetOrder() ([]*string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	data, err := r.ReadTodosData()
	if err != nil {
		return nil, err
	}

	return data.Order, nil
}
