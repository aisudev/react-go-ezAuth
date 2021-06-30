import React from 'react'
import Button from '../components/Button'
import { useController } from '../controllers/Controller'

export default function HomeScreen() {

    const controller = useController()

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
                my-5
                '
            >
                Hi, {controller.user.name}
            </span>
            <Button
                title='Log out'
                onClick={() => controller.SignOut()}
            />
        </div>
    )
}
