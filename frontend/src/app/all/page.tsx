import Image from 'next/image'

export default function ViewAll() {
  let nfts = ["1", "2", "3"]
  // make request to backend and get list of all nfts paginated


  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div>
        Casper generate NFT
      </div>
      <div className="flex flex-col gap-3">
        <div>Your NFTs</div>
        {nfts.map(e => {
          return (<div>{e}</div>)
        })}
      </div>
      <button>Make new</button>
      <a href="/">Go back</a>
    </main>
  )
}

