//entry point to import all elements to be bundle
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap/dist/js/bootstrap.bundle.js'
import ReactDOM from 'react-dom'
import React from 'react'
import './test.jsx'
import Test from './test.jsx'

ReactDOM.render(
    <div>
        <h1>Test Component</h1>
        <Test></Test>
    </div>,
    document.getElementById('root')
)
