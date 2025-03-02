import React, { useState, useEffect } from "react";
import { createUserAndStartGame, fetchNextQuestion, validateAnswer } from "../api/gameApi";
import ConfettiAnimation from "./Confetti";
import "bootstrap/dist/css/bootstrap.min.css";

const QuizGame = () => {
  const [userName, setUserName] = useState(null);
  const [sessionId, setSessionId] = useState(null);
  const [question, setQuestion] = useState(null);
  const [showConfetti, setShowConfetti] = useState(false);
  const [feedback, setFeedback] = useState(null);

  const startGame = async () => {
    const gameData = await createUserAndStartGame();
    if (gameData) {
      setUserName(gameData.userName); // Use backend-generated username
      setSessionId(gameData.sessionId);
      loadNextQuestion(gameData.userName, gameData.sessionId);
    }
  };

  const loadNextQuestion = async (user, session) => {
    const data = await fetchNextQuestion(user, session);
    if (data) {
      setQuestion(data);
      setFeedback(null);
    }
  };

  const handleAnswer = async (option) => {
    const result = await validateAnswer(userName, sessionId, question.clues[0].id, option.id);
    if (result.correct) {
      setShowConfetti(true);
      setFeedback("✅ Correct!");
      setTimeout(() => {
        setShowConfetti(false);
        loadNextQuestion(userName, sessionId);
      }, 2000);
    } else {
      setFeedback("❌ Wrong Answer!");
      setTimeout(() => loadNextQuestion(userName, sessionId), 1000);
    }
  };

  return (
    <div className="container text-center mt-5">
      <ConfettiAnimation show={showConfetti} />
      <h1 className="mb-4">Guess the City!</h1>

      {!userName ? (
        <button className="btn btn-primary" onClick={startGame}>
          Start Game
        </button>
      ) : (
        <div>
          {question && (
            <>
              <h4>Clues:</h4>
              <ul className="list-group">
                {question.clues.map((clue) => (
                  <li key={clue.clue_id} className="list-group-item">
                    {clue.clue}
                  </li>
                ))}
              </ul>

              <h4 className="mt-4">Options:</h4>
              <div className="row">
                {question.options.map((option) => (
                  <div key={option.id} className="col-md-3">
                    <button
                      className="btn btn-outline-primary btn-block mt-2"
                      onClick={() => handleAnswer(option)}
                    >
                      {option.city}, {option.country}
                    </button>
                  </div>
                ))}
              </div>

              {feedback && <h3 className="mt-4">{feedback}</h3>}

              <button
                className="btn btn-warning mt-4"
                onClick={() => loadNextQuestion(userName, sessionId)}
              >
                Skip
              </button>
            </>
          )}
        </div>
      )}
    </div>
  );
};

export default QuizGame;
