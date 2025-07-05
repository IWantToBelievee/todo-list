import './App.css'
import { DeleteToDo, SaveOrder } from '../wailsjs/go/main/App';
import SortableItem from './SortableItem';
import { DndContext } from '@dnd-kit/core';
import { arrayMove, SortableContext } from '@dnd-kit/sortable';
import {
  restrictToWindowEdges,
  createSnapModifier,
} from '@dnd-kit/modifiers'

export default function ToDoList({ todos, loadTodos, updateTodo, toggleCompleted, order, setOrder, loadOrder }) {
    const deleteTodo = async (id: string) => {
        await DeleteToDo(id);
        loadTodos();
        loadOrder();
    };

    const onDragEnd = (e) => {
        const {active, over} = e;
        if (over && active.id !== over.id) {
            setOrder((items) => {
                const oldIndex = items.findIndex((item) => item === active.id);
                const newIndex = items.findIndex((item) => item === over.id);
                const newOrder = arrayMove(items, oldIndex, newIndex);
                SaveOrder(newOrder);
                return newOrder;
            });
        }
    }

    return (
        <div className="flex-12 w-full ml-25">
            <div className='flex flex-row gap-4 flex-wrap pt-4'>
                <DndContext onDragEnd={onDragEnd} modifiers={[restrictToWindowEdges]}>
                    <SortableContext items={order}>
                        {order && order.map((item) => (
                            <SortableItem key={item} id={item} todo={todos.find(i => i.id === item)} deleteTodo={deleteTodo} updateTodo={updateTodo} toggleCompleted={toggleCompleted} />
                        ))}
                    </SortableContext>
                </DndContext>
            </div>
        </div>
    );
}