import {useState} from 'react';
import Context from '../generated/core/context';
import styled from 'styled-components';
import TodoNote from '../generated/api/clients/todo-backend/models/TodoNote';
import EntryItem from './EntryItem';
import TodoNoteForm, {UpdateListFunc} from './TodoNoteForm';
import {useAppSelector} from '../app/hooks';
import Button from './Button';
import TodoNoteUpdateRequest from '../generated/api/clients/todo-backend/models/TodoNoteUpdateRequest';
import moment from 'moment';

export interface TodoNoteEntryProps {
    content: TodoNote;
    updateList: UpdateListFunc
}

export default function TodoNoteEntry(props: TodoNoteEntryProps) {
    var [isEdit, setEdit] = useState(false);
    const note = props.content
    const updateList = (ctx: Context) => {
        props.updateList(ctx);
        setEdit(false);
    }

    var ctx = useAppSelector((rootState) => rootState.mifyState.value)
    const deleteNote = async () => {
        await ctx.clients.todoBackend().todosIdDelete(note.id)
            .catch(_ => alert("Failed to delete note"))
        props.updateList(ctx);
    }

    const toggleCompleteNote = async () => {
        const req = new TodoNoteUpdateRequest(note.title)
        req.description = note.description
        req.is_completed = !note.is_completed

        await ctx.clients.todoBackend().todosIdPut(note.id, req)
            .catch(_ => alert("Failed to delete note"))
        props.updateList(ctx);
    }

    return (
        <EntryItem>
            {!isEdit && <EntryData>
                <Content>
                    <Title>{note.title}</Title>
                    <Description>{note.description}</Description>
                    <Date>{moment(note.updated_at).fromNow()}</Date>
                </Content>
                <Buttons>
                    <CompleteButton
                        isCompleted={note.is_completed}
                        onClick={toggleCompleteNote}>
                        {note.is_completed ? "Uncomplete" : "Complete"}
                    </CompleteButton>
                    <Button onClick={() => setEdit(true)}>Edit</Button>
                    <DeleteButton onClick={deleteNote}>Delete</DeleteButton>
                </Buttons>
            </EntryData>}
            {isEdit && <TodoNoteForm note={note} updateList={updateList} />}
        </EntryItem>)

}

const EntryData = styled.div`
    display: flex;
`;

const Content = styled.div`
    flex: 1;
`;

const Title = styled.h2``;

const Description = styled.div`
    font-size: 16px;
    margin: 5px 0px;
`;

const Date = styled.div``;

const Buttons = styled.div`
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;

    button {
        margin-right: 5px;
    }
`;

interface CompleteButtonProps {
    isCompleted: boolean;
}

const CompleteButton = styled(Button)<CompleteButtonProps>`
    background-color: #${props => props.isCompleted ? "006400" : "8b8000"};
`;

const DeleteButton = styled(Button)`
    background-color: #8b0000;
`;
