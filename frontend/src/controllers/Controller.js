import { useState, useContext, createContext, useEffect } from 'react'
import { useHistory } from 'react-router-dom'
import API from './API'

const context = createContext()

export const useController = () => new ControllerClass(useContext(context))

class ControllerClass {

    constructor(context) {
        this.context = context
        this.username = context.username
        this.name = context.name
        this.password = context.password
        this.user = context.user
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

    SignUp() {
        this.context.SignUp()
    }

    SignIn() {
        this.context.SignIn()
    }

    SignOut() {
        this.context.SignOut()
    }

}

function Controller({ children }) {

    const [user, setUser] = useState({})
    const [username, setUsername] = useState('')
    const [name, setName] = useState('')
    const [password, setPassword] = useState('')

    const history = useHistory()

    useEffect(() => {
        API.GetUser()
            .then(resp => {

                setUser(resp.data.data)


            }).catch(err => {
                console.log(err)
                API.RefreshAuthentation()
                    .then(resp => {
                        const { access_token, refresh_token } = resp.data.data

                        localStorage.setItem('access_token', access_token)
                        localStorage.setItem('refresh_token', refresh_token)

                        window.location.reload()

                    }).catch(err => {
                        console.log('Refresh Auth:', err)
                    })
            })
    }, [])

    const SignUp = () => {

        const data = {
            username: username,
            password: password
        }

        API.CreateAuthentication(data)
            .then(resp => {

                API.GetAuthentication(data)
                    .then(resp => {
                        const { access_token, refresh_token } = resp.data.data

                        localStorage.setItem('access_token', access_token)
                        localStorage.setItem('refresh_token', refresh_token)

                        API.CreateUser({ name: name })
                            .then(resp => {
                                console.log('Done.')
                                history.push('/signin')

                            })
                            .catch(err => {
                                console.log(err)
                            })

                    }).catch(err => {
                        console.log(err)
                    })
            }).catch(err => {
                console.log(err)
            })

    }

    function SignIn() {

        const data = {
            username: username,
            password: password
        }

        API.GetAuthentication(data)
            .then(resp => {
                const { access_token, refresh_token } = resp.data.data

                localStorage.setItem('access_token', access_token)
                localStorage.setItem('refresh_token', refresh_token)

                API.GetUser()
                    .then(resp => {
                        setUser(resp.data.data)

                    }).catch(err => {
                        console.log(err)

                    })

            }).catch(err => {
                console.log(err)
            })
    }

    function SignOut() {
        setUser(null)
        localStorage.removeItem('access_token')
        localStorage.removeItem('refresh_token')
    }

    return (
        <context.Provider
            value={{
                username, setUsername,
                name, setName,
                password, setPassword,
                SignUp, SignIn, SignOut,
                user
            }}
        >
            {children}
        </context.Provider>
    )
}

export default Controller

