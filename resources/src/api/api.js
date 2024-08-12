import axios from 'axios';

const api = axios.create({
    baseURL: process.env.REACT_APP_API_URL || 'http://localhost:8080/api',
});

export const getExampleData = async () => {
    try {
        const response = await api.get('/example');
        return response.data;
    } catch (error) {
        console.error('Error fetching data:', error);
        throw error;
    }
};
