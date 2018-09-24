import React from 'react'
import {HashRouter} from 'react-router-dom'
import Routes from './routes'
import Nav from './main/nav'

export default class App extends React.Component {
    
    render(){
        return(
            <HashRouter>
                <div className="container">
                    <Nav />
                    <div className="row">
                        <Routes />
                    </div>
                </div>
            </HashRouter>
        )
    }
} 