import { useEffect, useState } from "react";
import KanbanBoard from "./components/KanbanBoard";
import { taskService } from "./services/api";

function App() {
  const [tasks, setTasks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    loadTasks();
  }, []);

  const loadTasks = async () => {
    try {
      setLoading(true);
      const data = await taskService.getAll();
      setTasks(data || []);
      setError(null);
    } catch (err) {
      setError("Erro ao carregar tarefas");
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const handleAddTask = async (title, description, status = "todo") => {
    try {
      const newTask = await taskService.create({ title, description, status });
      setTasks((current) => [newTask, ...current]);
    } catch (err) {
      setError("Erro ao adicionar tarefa");
    }
  };

  const handleUpdateTask = async (id, updatedTask) => {
    try {
      const updated = await taskService.update(id, updatedTask);
      setTasks((current) => current.map((task) => (task.id === id ? updated : task)));
    } catch (err) {
      setError("Erro ao atualizar tarefa");
    }
  };

  const handleDeleteTask = async (id) => {
    try {
      await taskService.delete(id);
      setTasks((current) => current.filter((task) => task.id !== id));
    } catch (err) {
      setError("Erro ao deletar tarefa");
    }
  };

  const handleMoveTask = async (taskId, newStatus) => {
    const task = tasks.find((t) => t.id === taskId);
    if (task) {
      try {
        const updated = await taskService.update(taskId, {
          ...task,
          status: newStatus,
        });
        setTasks((current) => current.map((t) => (t.id === taskId ? updated : t)));
      } catch (err) {
        setError("Erro ao mover tarefa");
      }
    }
  };

  return (
    <div className="w-full max-w-6xl mx-auto px-4 py-8 sm:px-6">
      <header className="mb-8 flex flex-col gap-4 rounded-2xl bg-slate-950 px-6 py-7 text-white shadow-xl sm:flex-row sm:items-center sm:justify-between sm:px-8">
        <div>
          <p className="mb-2 text-sm font-semibold uppercase tracking-[0.18em] text-indigo-300">Organize o seu trabalho</p>
          <h1 className="text-3xl font-bold tracking-tight sm:text-4xl">
            Veritas Mini Kanban
          </h1>
          <p className="mt-2 text-slate-300">
            Um fluxo simples, focado e persistente para suas tarefas.
          </p>
        </div>
        <div className="rounded-xl border border-white/10 bg-white/10 px-4 py-3 text-sm text-slate-200">
          {tasks.length} {tasks.length === 1 ? "tarefa no quadro" : "tarefas no quadro"}
        </div>
      </header>

      {error && (
        <div className="bg-red-500/90 text-white px-6 py-4 rounded-lg mb-6 flex justify-between items-center shadow-lg">
          <span>{error}</span>
          <button
            onClick={() => setError(null)}
            className="hover:bg-red-600 rounded px-2 py-1"
          >
            ✕
          </button>
        </div>
      )}

      {loading ? (
        <div className="flex flex-col items-center justify-center min-h-96 gap-5">
          <div className="w-12 h-12 border-4 border-white/30 border-t-white rounded-full animate-spin"></div>
          <p className="text-white text-lg font-medium">
            Carregando tarefas...
          </p>
        </div>
      ) : (
        <KanbanBoard
          tasks={tasks}
          onAddTask={handleAddTask}
          onUpdateTask={handleUpdateTask}
          onDeleteTask={handleDeleteTask}
          onMoveTask={handleMoveTask}
        />
      )}
    </div>
  );
}

export default App;
