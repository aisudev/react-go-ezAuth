import React from 'react'

export default function Input({
    placeHolder = '',
    type = 'text',
    onChange = () => { }
}) {
    return (
        <input type={type} placeholder={placeHolder} onChange={onChange}
            className='
                w-full
                px-4 py-2
                rounded-md
                bg-gray-100
                border-transparent
                focus:outline-none
                focus:ring-2 focus:ring-yellow-400
                focus:bg-white
            '
        />
    )
}
