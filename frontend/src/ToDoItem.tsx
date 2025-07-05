import { useEffect, useRef, useState } from 'react';
import { RxDragHandleDots2 } from "react-icons/rx";
import { FaRegTrashCan } from "react-icons/fa6";
import TextareaAutosize from 'react-textarea-autosize';
import { UpdateToDo } from '../wailsjs/go/main/App';


export default function ToDoItem({ todo, deleteTodo, updateTodo, toggleCompleted, listeners }) {
    const titleRef = useRef<HTMLTextAreaElement>(null);
    const [updateTimer, setUpdateTimer] = useState<NodeJS.Timeout>(null);

    const handleTitleChanged = async () => {
        if (titleRef.current) {
            titleRef.current.value = titleRef.current.value.replace(/\n/g, '');
            if (!updateTimer) {
                let timer = setTimeout(() => {
                    updateTodo({
                        id: todo.id,
                        title: titleRef.current.value.length > 0 ? titleRef.current.value : "Undefined" 
                    });
                    setUpdateTimer(null);
                }, 1000)
                setUpdateTimer(timer);
            }
        }
    }

    return (
        <div className={`card shadow-2xl rounded-2xl transition-opacity ${todo.completed ? 'opacity-50' : 'opacity-100'}`}>
            <div className="card-body flex flex-col justify-between min-w-70 max-w-70 min-h-50">
                <TextareaAutosize 
                    className="card-title focus:outline-none max-w-60 text-wrap resize-none" 
                    ref={titleRef} 
                    defaultValue={todo.title}
                    placeholder="Title"
                    onChange={() => handleTitleChanged()}
                />
                <div className='card-actions flex items-center justify-between'>
                    <input
                        className='checkbox' 
                        type='checkbox'
                        checked={todo.completed}
                        onClick={() => toggleCompleted(todo.id)}
                    />
                    <button
                        className='btn btn-ghost'
                        onClick={() => deleteTodo(todo.id)}
                    >
                        <FaRegTrashCan className='w-4 h-4'/>
                    </button>
                    <button
                        className='btn btn-ghost btn-circle'
                        {...listeners}
                    >
                        <RxDragHandleDots2 className='w-6 h-6' />
                    </button>
                </div>
            </div>
        </div>
    )
}