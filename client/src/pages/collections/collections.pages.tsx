import React from 'react';
import './collections.style.css';
import { CollectionSwiper, Topartist } from '../../components';
function Collections() {
  
  return (
    <div className='collection-page'>
      <h1>
        Collections
      </h1>
      <CollectionSwiper/>
      <div className="artist-container-section">
        <Topartist/>
      </div>
    </div>
    
  )
}

export default Collections;