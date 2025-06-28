import './App.css'
import { DeleteToDo } from '../wailsjs/go/main/App';
import ToDoItem from "./ToDoItem";
import ToDo from './types';

interface ToDoListProps {
    todos: ToDo[];
    loadTodos: Function;
}

export default function ToDoList({ todos, loadTodos }: ToDoListProps) {
    const deleteTodo = async (id: string) => {
        await DeleteToDo(id);
        loadTodos();
    };

    return (
        <div className="flex-12 w-full ml-25 transition-all delay-75">
            <div className='flex flex-row gap-4 flex-wrap pt-4'>
                {todos && todos.map((todo) => (
                    <ToDoItem todo={todo} deleteTodo={deleteTodo} />
                ))}
            </div>
        </div>
    );
}