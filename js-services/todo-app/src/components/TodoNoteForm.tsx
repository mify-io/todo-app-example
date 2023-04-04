import {useForm} from 'react-hook-form';
import {useAppSelector} from '../app/hooks';
import Context from '../generated/core/context';
import TodoNote from '../generated/api/clients/todo-backend/models/TodoNote';
import TodoNoteCreateRequest from '../generated/api/clients/todo-backend/models/TodoNoteCreateRequest';
import TodoNoteUpdateRequest from '../generated/api/clients/todo-backend/models/TodoNoteUpdateRequest';
import styled from 'styled-components';
import Button from './Button';

export type UpdateListFunc = (ctx: Context) => void;

type FormValues = {
    title: string;
    description: string;
    serverError: string;
};

interface TodoNoteFormProps {
    note?: TodoNote;
    updateList: UpdateListFunc
}


export default function TodoNoteForm(props: TodoNoteFormProps) {
    var ctx = useAppSelector((rootState) => rootState.mifyState.value)

    let defaultValues = {}
    if (props.note) {
        defaultValues = {
            title: props.note.title,
            description: props.note.description,
        }
    }

    const {register, handleSubmit, setError, formState: {errors}} = useForm<FormValues>({
        defaultValues: defaultValues,
    });

    async function onSubmit(data: FormValues) {
        if (props.note) {
            let req = TodoNoteUpdateRequest.constructFromObject(data, null);
            req.is_completed = props.note.is_completed;
            await ctx.clients.todoBackend().todosIdPut(props.note.id, req)
                .catch(_ => setError('serverError',
                    {type: "server", message: "Failed to update note"}))
            props.updateList(ctx)
            return;
        }
        const req = TodoNoteCreateRequest.constructFromObject(data, null);
        await ctx.clients.todoBackend().todosPost(req)
            .catch(_ => setError('serverError',
                {type: "server", message: "Failed to add note"}))
        props.updateList(ctx)
    }

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <div>{errors.serverError && errors.serverError.message}</div>
            <TitleInput placeholder="Title" type="text" {...register("title")} />
            <DescriptionInput
                placeholder="What to do you want to do today?"
                {...register("description")}
            />
            <Button type="submit">{props.note ? "Update" : "Add"} note</Button>
        </form>
    );
}

const TitleInput = styled.input`
    display: block;
    font-size: 18px;
    background-color: #282c34;
    color: white;
    border: none;
    border-radius: 2px;
    padding: 5px;
    margin-top: 10px;
    width: 100%;
`;

const DescriptionInput = styled.textarea`
    display: block;
    background-color: #282c34;
    color: white;
    border: none;
    border-radius: 2px;
    padding: 5px;
    margin-top: 10px;
    min-height: 100px;
    width: 100%;
`;
