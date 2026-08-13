import axios from "axios";

const API_BASE_URL = import.meta.env.VITE_API_URL || "http://localhost:8080/api";

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

export const taskService = {
  // Obter todas as tarefas
  getAll: async () => {
    const response = await api.get("/tasks");
    return response.data || [];
  },

  // Criar uma nova tarefa
  create: async (task) => {
    try {
      const response = await api.post("/tasks", {
        title: task.title,
        description: task.description,
        status: task.status || "todo",
      });
      return response.data;
    } catch (error) {
      console.error("Erro ao criar tarefa:", error);
      throw error;
    }
  },

  // Atualizar uma tarefa
  update: async (id, task) => {
    try {
      const response = await api.put(`/tasks/${id}`, {
        title: task.title,
        description: task.description,
        status: task.status,
      });
      return response.data;
    } catch (error) {
      console.error("Erro ao atualizar tarefa:", error);
      throw error;
    }
  },

  // Deletar uma tarefa
  delete: async (id) => {
    try {
      await api.delete(`/tasks/${id}`);
      return true;
    } catch (error) {
      console.error("Erro ao deletar tarefa:", error);
      throw error;
    }
  },

  // Obter uma tarefa específica
  getById: async (id) => {
    try {
      const response = await api.get(`/tasks/${id}`);
      return response.data;
    } catch (error) {
      console.error("Erro ao buscar tarefa:", error);
      return null;
    }
  },
};

export default api;
