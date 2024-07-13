const searchInput = document.querySelector(".search-wrapper input")
const artistList = document.querySelector(".main-container");

let artists = []
var locations = []
// locations[1] is an array of locations for artist with id 1

GetArtists(); // display all artists & fetch locations

searchInput.addEventListener("input", (e) => {
  const input = e.target.value.toLowerCase()
  const filteredArtists = artists.filter(artist =>
    artist.name.toLowerCase().includes(input) ||
    artist.members.some(member => member.toLowerCase().includes(input)) ||
    artist.firstAlbum.startsWith(input) ||
    artist.creationDate == input ||
    Math.floor(artist.creationDate/100) == input ||
    Math.floor(artist.creationDate/1000) == input ||
    Math.floor(artist.creationDate/10) == input ||
    isLocation(input, artist.id)
  );
  displayArtists(filteredArtists);
});

async function GetArtists() {
  try {
    const response = await fetch('https://groupietrackers.herokuapp.com/api/artists');
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    const data = await response.json();
    artists = data;

    // Fetch location data for each artist
    await Promise.all(artists.map(async (artist) => {
      await GetLocations(artist.id);
    }));

    displayArtists(artists);
  } catch (e) {
    console.error('Error in fetch request:', e);
  }
}

// add "https://cors-anywhere.herokuapp.com/" to solve CORS Protocol error
async function GetLocations(id) {
  try {
    const response = await fetch(`https://groupietrackers.herokuapp.com/api/locations/${id}`);
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    const data = await response.json();
    // location based on id
    locations[id] = data.locations;
  } catch (e) {
    console.error('Error in fetch request:', e);
  }
}

function isLocation(input, artistId) {
  for(let i=0; i<locations.length; i++){
    if(i == artistId){
      for(let j=0; j<locations[i].length; j++){
        if(locations[artistId].some(location => location.toLowerCase().includes(input))){
          return true
        }
      }
    }
  }
  return false
}
// OR
// if (locations[artistId]) {
//   return locations[artistId].some(location => location.toLowerCase().includes(input.toLowerCase()));
// }
// return false;


function displayArtists(artistData) {
  artistList.innerHTML = '';

  artistData.forEach(artist => {
    const artistCard = document.createElement('div');
    artistCard.classList.add('artist-card');

    const artistLink = document.createElement('a');
    artistLink.classList.add('ArtLink');
    artistLink.href = `/artist/${artist.id}`;

    const artistImage = document.createElement('img');
    artistImage.src = artist.image;
    artistImage.alt = 'Click here to go to the Artist page';

    const artistName = document.createElement('p');
    artistName.classList.add('artistName');
    artistName.textContent = artist.name;

    artistLink.appendChild(artistImage);
    artistLink.appendChild(artistName);
    artistCard.appendChild(artistLink);
    artistList.appendChild(artistCard);
  });
}