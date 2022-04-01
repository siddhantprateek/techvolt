import React, { useState } from 'react'
import { useEffect } from 'react'
import axios from 'axios'
import './Topartist.styles.css'
function Topartist() {
  const [artists, setArtists] = useState(null)
  useEffect(()=> {
    axios.get('http://localhost:4000/api/artist')
    .then((res) => {
      console.log(res.data.data)
      setArtists(res.data.data)
    })
    .catch((err) => console.error(err))
  }, [])

  const slicesArtist = artists?.slice(0, 10)
  return (
    <div className='top-artist-container'>
      <h2>Top Artist</h2>
      {slicesArtist?.map((artist, key) =>(
        <div className="artist-section">
          <div className="artist-logo">
            <h1>{key + 1}</h1>
            <img src={artist.profileImageUrl} alt="" />
          </div>
          <div className="artist-details">
            <h2 className='artist-name'>{artist.name}</h2>
            <p>@{artist.profileUrl}</p>
          </div>
        </div>
      ))}
    </div>
  )
}

export default Topartist