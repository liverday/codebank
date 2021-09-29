import axios from 'axios';

export default axios.create({
    baseURL: 'http://app:3000/api'
})