import React, { useEffect, useState } from 'react'
import ReactDOM from 'react-dom/client'

const ChildComponent = () => {
  useEffect(() => {
    console.log('Child Component Created')
    const timer = setInterval(() => console.log('...tick...'), 1000)

    // Cleanup function (ComponentWillUnmount)
    return () => {
      console.log('Component Unmounted (Cleanup ran)')
      clearInterval(timer)
    }
  }, [])

  return <p>I am alive! (Check console)</p>
}

const UnmountingExample = () => {
  const [show, setShow] = useState(true)
  return (
    <div style={{ padding: '20px', fontFamily: 'sans-serif' }}>
      <h1>Lifecycle: Unmounting</h1>
      <button onClick={() => setShow(!show)}>
        {show ? 'Unmount Child' : 'Mount Child'}
      </button>
      {show && <ChildComponent />}
    </div>
  )
}

ReactDOM.createRoot(document.getElementById('root')).render(<UnmountingExample />)