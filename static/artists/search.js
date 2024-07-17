const searchInput = document.querySelector(".search-wrapper input");
const artistList = document.querySelector(".main-container");

let artists = [];
let locations = []; // locations[1] means locations for artist with id=1

GetArtists(); // Fetch artists and display them

searchInput.addEventListener("input", (e) => {
  const input = e.target.value.toLowerCase();

  if (input === "") {
    clearSuggestions();
    displayArtists(artists); // Optionally, you can display all artists
    return;
  }

  const filteredArtists = [];
  const suggestions = [];

  artists.forEach(artist => {
    let matched = false;

    // Check if the artist's name matches the input
    if (artist.name.toLowerCase().includes(input)) {
      filteredArtists.push(artist);
      suggestions.push({text: `${artist.name} - artist/band`, id:artist.id});
      matched = true;
    }

    // Check if any member's name matches the input
    artist.members.forEach(member => {
      if (member.toLowerCase().includes(input)) {
        if (!matched) {
          filteredArtists.push(artist);
          matched = true;
        }
        suggestions.push({text: `${member} - member in ${artist.name}`, id:artist.id});
      }
    });

    // Check if the artist's first album matches the input
    if (artist.firstAlbum.includes(input)) {
      if (!matched) {
        filteredArtists.push(artist);
        matched = true;
      }
      suggestions.push({text:`${artist.firstAlbum} - first album for ${artist.name}`, id:artist.id});
    }

    // Check if the artist's creation date matches the input
    if (artist.creationDate.toString().includes(input)) {
      if (!matched) {
        filteredArtists.push(artist);
        matched = true;
      }
      suggestions.push({text:`${artist.creationDate} - creation date for ${artist.name}`, id:artist.id});
    }

    // Check if the input matches any locations of the band
    if (locations[artist.id]) { // if this artist has location
      locations[artist.id].forEach(location => {
        if (location.toLowerCase().includes(input)) {
          suggestions.push({text:`${location} - location in ${artist.name}`, id:artist.id});
          if (!matched) {
            filteredArtists.push(artist);
            matched = true;
          }
        }
      });
    }
  });

  if (filteredArtists.length === 0) {
    clearSuggestions();
    displayNoResult();
  } else {
    clearSuggestions();
    displaySuggestions(suggestions);
    displayArtists(filteredArtists);
  }
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
    await Promise.all(artists.map(artist => GetLocations(artist.id)));

    displayArtists(artists);
  } catch (e) {
    console.error('Error in fetch request:', e);
  }
}
async function GetLocations(id) {
  try {
    const response = await fetch(`https://groupietrackers.herokuapp.com/api/locations/${id}`);
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    const data = await response.json();
    locations[id] = data.locations;
  } catch (e) {
    console.error('Error in fetch request:', e);
  }
}

// function isLocation(input, artistId) {
//   if (locations[artistId]) {
//     return locations[artistId].some(location => location.toLowerCase().includes(input.toLowerCase()));
//   }
//   return false;
//   // OR
// //   for(let i=0; i<locations.length; i++){
// //     if(i == artistId){
// //       for(let j=0; j<locations[i].length; j++){
// //         if(locations[artistId].some(location => location.toLowerCase().includes(input))){
// //           return true
// //         }
// //       }
// //     }
// //   }
// //   return false
// // }
// }

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

function displayNoResult() {
  artistList.innerHTML = '<p>No results found</p>';
}

function displaySuggestions(suggestions) {
  const resBox = document.querySelector('.result-box');
  resBox.innerHTML = ''; // Clear previous suggestions

  const uList = document.createElement('ul');
  resBox.appendChild(uList);

  suggestions.forEach(suggestion => {
    const suggestionElement = document.createElement('li');
    let linkTo = document.createElement('a')
    suggestionElement.appendChild(linkTo)
    linkTo.textContent = suggestion.text;
    const [name, type] = suggestion.text.split(" - ")
    suggestionElement.addEventListener('click', () => {
      searchInput.value = name;
      linkTo.href = `http://localhost:8080/artist/${suggestion.id}`
      window.location.href= linkTo.href;
      resBox.innerHTML = ''; // Clear suggestions
    });

    uList.appendChild(suggestionElement);
  });
}

function clearSuggestions() {
  const resBox = document.querySelector('.result-box');
  resBox.innerHTML = '';
}
