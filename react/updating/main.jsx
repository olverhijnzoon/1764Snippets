import React, { useEffect, useState } from 'react'
import ReactDOM from 'react-dom/client'

const UpdatingExample = () => {
  const [count, setCount] = useState(0)

  useEffect(() => {
    console.log(`Component Updated: Count is ${count}`)
  }, [count]) // Runs whenever 'count' changes

  return (
    <div style={{ padding: '20px', fontFamily: 'sans-serif' }}>
      <h1>Lifecycle: Updating</h1>
      <button onClick={() => setCount(c => c + 1)}>Count: {count}</button>
    </div>
  )
}

ReactDOM.createRoot(document.getElementById('root')).render(<UpdatingExample />)