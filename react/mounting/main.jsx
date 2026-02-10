import React, { useEffect, useState } from 'react'
import ReactDOM from 'react-dom/client'

const MountingExample = () => {
  const [status, setStatus] = useState('Initializing...')

  useEffect(() => {
    console.log('Component Mounted (This runs ONCE)')
    
    // Simulate a delay, like fetching data
    setTimeout(() => {
      setStatus('Mounted & Ready!')
    }, 1000)

  }, []) // Empty dependency array = ComponentDidMount

  return (
    <div style={{ padding: '20px', fontFamily: 'sans-serif' }}>
      <h1>Lifecycle: Mounting</h1>
      <p>Status: <strong>{status}</strong></p>
      <small>Check the browser console logs.</small>
    </div>
  )
}

ReactDOM.createRoot(document.getElementById('root')).render(<MountingExample />)