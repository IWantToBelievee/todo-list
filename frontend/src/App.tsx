import ToDoList from './ToDoList';
import { HiPlus } from 'react-icons/hi';
import { CreateToDo, GetOrder, GetToDoByID, GetToDos, UpdateToDo } from '../wailsjs/go/main/App';
import { PiPaintBrushHousehold } from "react-icons/pi";
import { useEffect, useState } from 'react';
import { themeChange } from "theme-change";
import { ToDo, UpdateRequest } from './types';

function App() {
    const [todos, setTodos] = useState<ToDo[]>([]);
    const [order, setOrder] = useState<string[]>([]);

    const loadTodos = async () => {
        try {
            setTodos([...(await GetToDos())])
        } catch (error) {
            setTodos([])
        }
    }

    const loadOneTodo = async (id: string) => {
        let todo = await GetToDoByID(id);
        setTodos((prev) => prev.map((i) => i.id == todo.id ? todo : i));
    }

    const updateTodo = (req: UpdateRequest) => {
        UpdateToDo(req);
        loadOneTodo(req.id);
    }

    const toggleCompleted = (id: string) => {
        UpdateToDo({id: id, completed: todos.find((t) => t.id == id).completed ? false : true});
        loadOneTodo(id);
    }

    const loadOrder = async () => {
        try {
            setOrder(await GetOrder());
        } catch (error) {
            setOrder([]);
        }
    }

    useEffect(() => {
        loadTodos();
    }, []);

    useEffect(() => {
        loadOrder();
    }, []);

    const handleCreateTodo = () => {
        CreateToDo();
        loadTodos();
        loadOrder();
    };

    const themes = [
        { value: "default", label: "Default" },
        { value: "retro", label: "Retro" },
        { value: "cyberpunk", label: "Cyberpunk" },
        { value: "valentine", label: "Valentine" },
        { value: "halloween", label: "Halloween" },
    ];

    useEffect(() => {
        themeChange(false);
    }, []);

    return (
        <div id="App">
            <div className="flex gap-6 p-0">
                <div 
                    className="flex flex-1 justify-center p-4 shadow-2xl rounded-r-2xl h-[100vh] fixed w-20"
                >
                    <div className='flex flex-col justify-between'>
                        <div className='flex flex-col gap-4 items-center'>
                            <button 
                                className="btn btn-ghost btn-circle"
                                onClick={() => handleCreateTodo()}
                            >
                                <HiPlus className='w-8 h-8'/>
                            </button>
                        </div>
                        <div className='flex flex-col gap-4 items-center'>
                            <div className="dropdown dropdown-top">
                                <div tabIndex={0} role="button" className="btn btn-ghost btn-circle m-1">
                                    <PiPaintBrushHousehold className='w-6 h-6' />
                                </div>
                                <ul tabIndex={0} className="dropdown-content bg-base-300 rounded-box z-[1] w-52 p-2 shadow-2xl">
                                    {themes.map((theme) => (
                                        <li key={theme.value}>
                                            <input
                                                type="radio"
                                                name="theme-dropdown"
                                                className="theme-controller w-full btn btn-sm btn-block btn-ghost justify-start"
                                                aria-label={theme.label}
                                                value={theme.value}
                                                data-set-theme={theme.value}
                                                data-act-class="btn-active"
                                            />
                                        </li>
                                    ))}
                                </ul>
                            </div>
                        </div>
                    </div>
                </div>
                {todos ? (<ToDoList 
                            todos={todos} 
                            loadTodos={loadTodos} 
                            updateTodo={updateTodo}
                            toggleCompleted={toggleCompleted}
                            order={order} 
                            setOrder={setOrder} 
                            loadOrder={loadOrder}
                        />) : (
                    <h2>There are no todos</h2>
                )}
            </div>
        </div>
    )
}

export default App
