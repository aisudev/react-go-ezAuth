import React from 'react'
import Input from '../components/Input'
import Button from '../components/Button'
import Title from '../components/Title'
import { useController } from '../controllers/Controller'

export default function SignUpScreen() {

    const controller = useController()

    return (
        <div
            className='
                container 
                mx-auto 
                md:px-32
                px-4
                my-52
            '
        >
            <div className='my-10' >
                <Title>What's World</Title>
            </div>
            <div
                className='
                    grid 
                    grid-cols-1 
                    gap-4
                '
            >
                <div>
                    <Input
                        placeHolder='enter your username'
                        type='text'
                        onChange={e => controller.setUsername(e.target.value)}
                    />
                </div>

                <div>
                    <Input
                        placeHolder='enter your name'
                        type='text'
                        onChange={e => controller.setName(e.target.value)}
                    />
                </div>

                <div>
                    <Input
                        placeHolder='enter your password'
                        type='password'
                        onChange={e => controller.setPassword(e.target.value)}
                    />
                </div>

                <Button
                    title='Sign Up'
                    onClick={() => controller.SignUp()}
                />
            </div>
        </div>
    )
}
