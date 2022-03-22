import React from 'react'
import './explore.styles.css'
import { Button } from '../../components'
import { NFTIMG1, NFTIMG2, NFTIMG3, NFTIMG4 } from '../../static'


function Explore() {
  return (
    <div className='explore-page'>
      <div className='hero-container'>
        <div className="hero-contents">
          <h1>Discover and Collect extraordianary NFTs</h1>
          <p>TechVolt is the NFT collections and marketplace.</p>
          <a href="/collections">

            <Button placeholder={"Collection"} color="blue"/>
          </a>
          <Button placeholder={"Create"} color="transparent"/>
        </div>
      </div>
      <div className="nft-container">
        <div className="contanier-1">
          <img className='nft-img' src={NFTIMG1} alt="" />
          <img className='nft-img' src={NFTIMG2} alt="" />
        </div>
        <div className="contanier-2">
          <img className='nft-img' src={NFTIMG3} alt="" />
          <img className='nft-img' src={NFTIMG4} alt="" />
        </div>
      </div>
    </div>
  )
}

export default Explore;