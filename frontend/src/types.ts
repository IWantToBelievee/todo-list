export interface ToDo {
    id: string;
    title: string;
    completed: boolean;
    created_at: Date;
}

export interface UpdateRequest {
    id: string;
    title: string | null;
}