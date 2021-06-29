import React from 'react'

export default function Button({
    title = 'click',
    type = 'submit',
    onClick = () => { }
}) {
    return (
        <button type={type} onClick={onClick}
            className='
            bg-yellow-400 
            hover:bg-yellow-500 
            focus:outline-none 
            focus:ring-2 
            focus:ring-yellow-600 
            focus:ring-opacity-50
            w-full
            px-4
            py-2
            shadow
            rounded-md
            text-white
            font-medium
            '
        >
            {title}
        </button>
    )
}
