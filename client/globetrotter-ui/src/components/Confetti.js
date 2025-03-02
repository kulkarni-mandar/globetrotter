import React from "react";
import Confetti from "react-confetti";
import { useWindowSize } from "react-use";

const ConfettiAnimation = ({ show }) => {
  const { width, height } = useWindowSize();
  return show ? <Confetti width={width} height={height} /> : null;
};

export default ConfettiAnimation;
