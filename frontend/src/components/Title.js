import React from 'react'

export default function Title({children}) {
    return (
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
            '
        >
            {children}
        </span>
    )
}
