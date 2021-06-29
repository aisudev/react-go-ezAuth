import React from 'react'
import Button from '../components/Button'

export default function HomeScreen() {
    return (
        <div
            className='
            container lg:px-96
            my-56 mx-auto
            '
        >
            <span
                className='
            text-5xl
            md:text-8xl
            text-center
            justify-center
            items-center
            font-extrabold 
            text-transparent 
            bg-clip-text
            bg-gradient-to-br from-yellow-400 to-red-600
            flex
            my-4
            '
            >
                Hi, User
            </span>
            <Button title='Log out' />
        </div>
    )
}
