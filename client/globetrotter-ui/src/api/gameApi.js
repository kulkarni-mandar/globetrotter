import axios from "axios";

const API_BASE_URL = "http://localhost:8080/api";

export const createUserAndStartGame = async () => {
  try {
    const userResponse = await axios.post(`${API_BASE_URL}/users`);
    const userName = userResponse.data.user_name; // Use username from backend

    const gameResponse = await axios.post(`${API_BASE_URL}/game/play/${userName}`);
    return { userName, sessionId: gameResponse.data.session_id };
  } catch (error) {
    console.error("Error starting game:", error);
    return null;
  }
};

export const fetchNextQuestion = async (userName, sessionId) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/game/play/${userName}/${sessionId}/next`);
    return response.data;
  } catch (error) {
    console.error("Error fetching next question:", error);
    return null;
  }
};

export const validateAnswer = async (userName, sessionId, clueId, selectedOptionId) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/game/play/${userName}/${sessionId}/${clueId}/validate`, {
      id: selectedOptionId,
    });
    return response.data;
  } catch (error) {
    console.error("Error validating answer:", error);
    return { correct: false };
  }
};

export const endGame = async (userName, sessionId) => {
    try {
      const response = await axios.post(`${API_BASE_URL}/game/play/${userName}/${sessionId}/end`);
      return response.data;
    } catch (error) {
      console.error("Error ending game:", error);
      return null;
    }
  };