import ToDoList from './ToDoList';
import { HiPlus } from 'react-icons/hi';
import { CreateToDo, GetToDos } from '../wailsjs/go/main/App';
import { useEffect, useState } from 'react';
import ToDo from './types';

function App() {
    const [todos, setTodos] = useState<ToDo[]>([]);
    const loadTodos = async () => {
        try {
            setTodos([...(await GetToDos())])
        } catch (error) {
            setTodos([])
        }
    }

    useEffect(() => {
        loadTodos();
    }, []);

    return (
        <div id="App">
            <div className="flex gap-6 p-0">
                <div 
                    className="flex flex-1 justify-center p-4 shadow-2xl rounded-r-2xl h-[100vh] fixed w-20"
                >
                    <button 
                        className="btn btn-ghost btn-circle"
                        onClick={async () => {
                            await CreateToDo("Undefined");
                            loadTodos();
                        }}
                    >
                        <HiPlus className='w-8 h-8'/>
                    </button>
                </div>
                {todos ? (<ToDoList todos={todos} loadTodos={loadTodos} />) : (
                    <h2>There are no todos</h2>
                )}
            </div>
        </div>
    )
}

export default App
