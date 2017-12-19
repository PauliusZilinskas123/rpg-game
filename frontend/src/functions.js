export function getURLParameter(name) {
    return decodeURIComponent((new RegExp('[?|&]' + name + '=' + '([^&;]+?)(&|#|;|$)').exec(location.search) || [null, ''])[1].replace(/\+/g, '%20')) || null;
}

export function getState() {
  return localStorage.getItem("state");
}

export function setState(state) {
  localStorage.setItem("state", state);
}

export function deleteState() {
  localStorage.removeItem("state");
}

export function getUser() {
  return localStorage.getItem("user");
}

export function setUser(user) {
  localStorage.setItem("user", user);
}

export function deleteUser() {
  localStorage.removeItem("user");
}