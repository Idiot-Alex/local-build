export default class ToolService {
  getTools() {
    return fetch('demo/data/tools.json')
        .then((res) => res.json())
        .then((d) => d.data);
  }
}