import { useRef, useState } from 'react';
import { FaRegTrashCan } from "react-icons/fa6";
import ToDo from './types';
import TextareaAutosize from 'react-textarea-autosize';
import { UpdateToDoCompleted, UpdateToDoTitle } from '../wailsjs/go/main/App';

interface ToDoOItemProps {
    todo: ToDo;
    deleteTodo: Function;
}

export default function ToDoItem({ todo, deleteTodo }: ToDoOItemProps) {
    const titleRef = useRef<HTMLTextAreaElement>(null);
    const [completed, setCompleted] = useState<boolean>(false);

    return (
        <div className={`card shadow-2xl rounded-2xl transition-opacity delay-75 ${completed ? 'opacity-50' : 'opacity-100'}`}>
            <div className="card-body flex flex-col justify-between min-w-70 max-w-70 min-h-50">
                <TextareaAutosize 
                    className="card-title focus:outline-none max-w-60 text-wrap resize-none" 
                    ref={titleRef} 
                    defaultValue={todo.title}
                    placeholder="Title"
                    onChange={async () => {
                        if (titleRef.current) {
                            titleRef.current.value = titleRef.current.value.replace(/\n/g, '');
                            await UpdateToDoTitle(todo.id, titleRef.current?.value.length! > 0 ? titleRef.current?.value : "Undefined");
                        }
                    }}
                />
                <div className='card-actions flex items-center justify-between'>
                    <input
                        className='checkbox' 
                        type='checkbox'
                        onClick={(e) => {
                            UpdateToDoCompleted(todo.id, e.currentTarget.checked);
                            setCompleted(e.currentTarget.checked);
                        }}
                    />
                    <button
                        className='btn btn-ghost'
                        onClick={() => deleteTodo(todo.id)}
                    >
                        <FaRegTrashCan className='w-4 h-4'/>
                    </button>
                </div>
            </div>
        </div>
    )
}