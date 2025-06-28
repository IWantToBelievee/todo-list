import './App.css'
import { useEffect, useState } from "react";
import { HiPlus } from "react-icons/hi";
import { GetToDos, CreateToDo, DeleteToDo, UpdateToDoTitle } from '../wailsjs/go/main/App';
import ToDoItem from "./ToDoItem";
import ToDo from './types';

export default function ToDoList() {
    const [todos, setTodos] = useState<ToDo[]>();
    const [sideOpen, setSideOpen] = useState<boolean>(false);

    const loadTodos = async () => {
        try {
            setTodos([...(await GetToDos())])
        } catch (error) {
            setTodos([])
            console.error("Error calling GetToDos:", error);
        }
    }

    const deleteTodo = async (id: string) => {
        await DeleteToDo(id);
        loadTodos();
    };

    useEffect(() => {
        loadTodos();
    }, []);

    return (
        <div className="flex gap-6 p-0">
            <div 
                className="flex flex-1 justify-center p-4 shadow-2xl rounded-r-2xl h-[100vh] fixed w-20"
                onMouseEnter={() => setSideOpen(true)}
                onMouseLeave={() => setSideOpen(false)}
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
            <div className="flex-12 w-full ml-25 transition-all delay-75">
                <div className='flex flex-row gap-4 flex-wrap pt-4'>
                    {todos && todos.map((todo) => (
                        <ToDoItem todo={todo} deleteTodo={deleteTodo} />
                    ))}
                </div>
            </div>
        </div>
    );
}