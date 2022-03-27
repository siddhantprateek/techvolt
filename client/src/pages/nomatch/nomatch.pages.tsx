
import { Link } from 'react-router-dom'
import "./nomatch.style.css"

function NoMatch() {
  return (
    <div className='no-match-page'>
      <div className="no-match-content">
        <h2>Are Your Lost</h2>
        <Link to="/">
          <button>Back To Home</button>
        </Link>
      </div>
    </div>
  )
}

export default NoMatch;