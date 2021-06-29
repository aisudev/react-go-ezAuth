import { useState, useContext, createContext } from 'react'

const context = createContext()

export const useController = () => new ControllerClass(useContext(context))

class ControllerClass {

    constructor(context) {
        this.context = context
        this.username = context.username
        this.name = context.name
        this.password = context.password
    }

    setUsername(data) {
        this.context.setUsername(data)
    }

    setName(data) {
        this.context.setName(data)
    }

    setPassword(data) {
        this.context.setPassword(data)
    }

}

function Controller({ children }) {

    const [username, setUsername] = useState('')
    const [name, setName] = useState('')
    const [password, setPassword] = useState('')

    return (
        <context.Provider
            value={{
                username, setUsername,
                name, setName,
                password, setPassword
            }}
        >
            {children}
        </context.Provider>
    )
}

export default Controller

