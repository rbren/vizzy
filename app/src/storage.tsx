export const getItem = function(key: string): string {
  try {
    window.localStorage;
  } catch (e) {
    return "";
  }
  return localStorage.getItem(key) || '';
}

export const setItem = function(key: string, value: string): void {
  try {
    window.localStorage;
  } catch (e) {
    return;
  }
  localStorage.setItem(key, value);
}

export const removeProject = function(key: string): void {
  try {
    window.localStorage;
  } catch (e) {
    return;
  }
  let projects = JSON.parse(getItem('projects') || '{}');
  delete projects[key];
  localStorage.setItem('projects', JSON.stringify(projects));
}
