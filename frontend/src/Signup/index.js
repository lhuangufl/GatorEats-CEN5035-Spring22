import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import NavBar from "./../NavBar/index";

export default function Signup() {
  const thisWidth = window.innerWidth;
  const thisHeight = window.innerHeight;

  const navigate = useNavigate();
  return (
    <div>
      <NavBar />
      <span>signup page</span>
    </div>
  );
}
