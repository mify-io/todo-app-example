import {useEffect, useState} from 'react';
import {useAppSelector} from './app/hooks';
import Context from './generated/core/context';
import TodoNote from './generated/api/clients/todo-backend/models/TodoNote';
import styled from 'styled-components';
import TodoNoteEntry from './components/TodoNoteEntry';
import TodoNoteForm from './components/TodoNoteForm';
import EntryItem from './components/EntryItem';

function App() {
    var [notes, setNotes] = useState([]);
    var ctx = useAppSelector((rootState) => rootState.mifyState.value)
    const updateList = async (ctx: Context) => {
        setNotes(await ctx.clients.todoBackend().todosGet())
    }
    useEffect(() => {updateList(ctx)}, [ctx])
    return (
        <div className="App">
            <Container>
                <h1>Mify To-Do List</h1>

                <EntryItem>
                    <TodoNoteForm updateList={updateList} />
                </EntryItem>
                {notes.map((note: TodoNote, idx) => <TodoNoteEntry key={idx} content={note} updateList={updateList} />)}
            </Container>
        </div>
    );
}

const Container = styled.div`
  position: relative;
  max-width: 800px;
  width: 100%;
  margin: 0 auto;
  margin-top: 30px;
  padding: 0 2rem;
`;

export default App;
