let isSubmitting = false;
// error Span
const errorSpan = document.querySelector(".error");

// signup form
const signUpForm = document.querySelector(".signup-form");

signUpForm &&
  signUpForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    if (isSubmitting) return;
    isSubmitting = true;
    const username = e.currentTarget.querySelector("#username").value;
    const email = e.currentTarget.querySelector("#email").value;
    const password = e.currentTarget.querySelector("#password").value;
    if (validateCredentials({ username, email, password }, "signup")) {
      sendSignUpInfo({ username, email, password });
    } else {
      isSubmitting = false;
    }
  });

const sendSignUpInfo = async (user) => {
  try {
    const data = await fetch("/api/signup", {
      method: "post",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(user),
    });
    if (data.ok) {
      window.location.href = "/login";
    } else {
      errorSpan.textContent = await data.text();
    }
  } catch (error) {
    errorSpan.textContent = "An error occurred. Please try again.";
  } finally {
    isSubmitting = false;
  }
};

// login form
const LoginForm = document.querySelector(".login-form");

LoginForm &&
  LoginForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    if (isSubmitting) return;
    isSubmitting = true;
    const email = e.currentTarget.querySelector("#email").value;
    const password = e.currentTarget.querySelector("#password").value;
    if (validateCredentials({ email, password }, "login")) {
      sendLoginInfo({ email, password });
    } else {
      isSubmitting = false;
    }
  });

const sendLoginInfo = async (user) => {
  try {
    const data = await fetch("/api/login", {
      method: "post",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(user),
    });
    if (data.ok) {
      window.location.href = "/";
    } else {
      errorSpan.textContent = await data.text();
    }
  } catch (error) {
    errorSpan.textContent = "An error occurred. Please try again.";
  } finally {
    isSubmitting = false;
  }
};

// helper
const errorMessages = {
  required: "Please enter all fields",
  emailLength: "Email must be between 5 and 50 characters.",
  passwordLength: "Password must be between 6 and 30 characters.",
  usernameLength: "Username must be between 3 and 15 characters.",
};
const validateCredentials = (user, action) => {
  const commonValidations = (fields) => {
    for (const field of fields) {
      if (!user[field]) {
        errorSpan.textContent = errorMessages.required;
        return false;
      }
    }
    return true;
  };

  if (action === "login") {
    const { email, password } = user;

    if (!commonValidations(["email", "password"])) return false;
    if (!validateEmail(email)) return false;
    if (!validatePassword(password)) return false;
  } else if (action === "signup") {
    const { username, email, password } = user;

    if (!commonValidations(["username", "email", "password"])) return false;
    if (!validateEmail(email)) return false;
    if (!validatePassword(password)) return false;

    if (username.length < 3 || username.length > 15) {
      errorSpan.textContent = errorMessages.usernameLength;
      return false;
    }
  }

  return true;
};

const validateEmail = (email) => {
  if (email.length < 5 || email.length > 50) {
    errorSpan.textContent = errorMessages.emailLength;
    return false;
  }
  return true;
};

const validatePassword = (password) => {
  if (password.length < 6 || password.length > 30) {
    errorSpan.textContent = errorMessages.passwordLength;
    return false;
  }
  return true;
};

// Get Info Data
export const getInfoData = async () => {
  const res = await fetch("/api/info");
  if (res.ok) {
    return await res.json();
  }
};
