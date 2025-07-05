import { CSS } from '@dnd-kit/utilities';
import ToDoItem from './ToDoItem';
import {useSortable} from '@dnd-kit/sortable';

export default function SortableItem({ id, todo, deleteTodo, updateTodo, toggleCompleted }) {
    const {
        attributes,
        listeners,
        setNodeRef,
        transform,
        transition,
    } = useSortable({id: id});

    const style = {
        transform: CSS.Transform.toString(transform),
        transition,
    };

    return (
        <div ref={setNodeRef} style={style} {...attributes}>
            <ToDoItem todo={todo} deleteTodo={deleteTodo} updateTodo={updateTodo} toggleCompleted={toggleCompleted} listeners={listeners}/>
        </div>
    );
}