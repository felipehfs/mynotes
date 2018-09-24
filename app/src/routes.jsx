import React from 'react'

import {Switch, Route, Redirect } from 'react-router'

import Note from "./notes/note"
import NoteDisplay from './notes/notedisplay'
import UpdateNote from './notes/updateNote'

export default props => 
<Switch>
    <Route exact path="/" component={Note} />
    <Route exact path="/notes" component={NoteDisplay} />
    <Route path="/notes/:id" component={UpdateNote} />
    <Redirect from="*" to="/" />
</Switch>