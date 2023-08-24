import axios from "axios"

const URL = "http://localhost:8080/user"

export const createUser = async (userData) => {
    try {
        const response = await axios.post(URL, userData);
        return response.data;
    } catch (error) {
        console.error('Error creating user:', error);
        throw error;
    }
};

export const deleteUserById = async (userId) => {
    try {
        const response = await axios.delete(`${URL}/${userId}`);
        return response.data;
    } catch (error) {
        console.error('Error deleting user:', error);
        throw error;
    }
};

export const getAllUsers = async () => {
    try {
        const response = await axios.get(`${URL}`);
        return response.data;
    } catch (error) {
        console.error('Error fetching users:', error);
        throw error;
    }
};

export const getUserById = async (userId) => {
    try {
        const response = await axios.get(`${URL}/${userId}`)
        return response.data
    } catch (error) {
        console.log(error)
        throw error
    }
}