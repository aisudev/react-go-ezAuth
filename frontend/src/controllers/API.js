import axios from 'axios'

const URL = `http://localhost:9999`

const CreateAuthentication = async (data = {
    username: '',
    password: ''
}) => {
    return await axios.post(`${URL}/auth`, data)
}

const GetAuthentication = async (data = {
    username: '',
    password: ''
}) => {
    return await axios.post(`${URL}/auth/signin`, data)
}

const RefreshAuthentation = async () => {
    const data = {
        refreshToken: localStorage.getItem('refresh_token')
    }

    return await axios.post(`${URL}/auth/refresh`, data)
}

const CreateUser = async (data = {
    name: ''
}) => {
    const accessToken = localStorage.getItem('access_token')

    return await axios.post(`${URL}/user`, data, {
        headers: {
            'Authorization': `Bearer ${accessToken}`
        }
    })
}

const GetUser = async () => {
    const accessToken = localStorage.getItem('access_token')

    return await axios.get(`${URL}/user`, {
        headers: {
            "Authorization": `Bearer ${accessToken}`
        }
    })
}

export default {
    CreateAuthentication,
    CreateUser,
    GetAuthentication,
    GetUser,
    RefreshAuthentation
}