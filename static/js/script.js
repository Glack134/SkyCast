async function searchWeather() {
    const city = document.getElementById('cityInput').value;
    if (!city) {
        alert('Please enter a city name');
        return;
    }

    try {
        const response = await fetch(`/main/search?city=${city}`);
        const data = await response.json();
        displayWeather(data.weather, 'weatherResult');
    } catch (error) {
        console.error('Error fetching weather data:', error);
        alert('Failed to fetch weather data');
    }
}

function displayWeather(weather, elementId) {
    const weatherResult = document.getElementById(elementId);
    if (!weather) {
        weatherResult.innerHTML = '<p>No weather data found.</p>';
        return;
    }

    const weatherCard = document.createElement('div');
    weatherCard.className = 'weather-card';
    weatherCard.innerHTML = `
        <h2>${weather.name}</h2>
        <p>Temperature: ${weather.main.temp}°C</p>
        <p>Condition: ${weather.weather[0].description}</p>
        <p>Humidity: ${weather.main.humidity}%</p>
    `;
    weatherResult.innerHTML = '';
    weatherResult.appendChild(weatherCard);
}

// Fetch weather for known cities on page load
async function fetchKnownCitiesWeather() {
    const cities = ['London', 'New York', 'Tokyo', 'Paris', 'Berlin'];
    const citiesWeather = document.getElementById('citiesWeather');
    citiesWeather.innerHTML = '';

    for (const city of cities) {
        try {
            const response = await fetch(`/main/search?city=${city}`);
            const data = await response.json();
            displayWeather(data.weather, 'citiesWeather');
        } catch (error) {
            console.error(`Error fetching weather data for ${city}:`, error);
        }
    }
}

// Call fetchKnownCitiesWeather on page load
window.onload = fetchKnownCitiesWeather;

