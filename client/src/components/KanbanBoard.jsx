import { useState } from "react";
import {
  FaClipboardList,
  FaTasks,
  FaCheckCircle,
  FaPlus,
} from "react-icons/fa";
import Column from "./Column";

const statuses = ["todo", "doing", "done"];
const statusLabels = {
  todo: "A Fazer",
  doing: "Em Progresso",
  done: "Concluído",
};

const statusIcons = {
  todo: FaClipboardList,
  doing: FaTasks,
  done: FaCheckCircle,
};

function KanbanBoard({
  tasks,
  onAddTask,
  onUpdateTask,
  onDeleteTask,
  onMoveTask,
}) {
  const [newTaskTitle, setNewTaskTitle] = useState("");
  const [newTaskDesc, setNewTaskDesc] = useState("");
  const [newTaskStatus, setNewTaskStatus] = useState("todo");

  const handleAddTask = (e) => {
    e.preventDefault();
    if (newTaskTitle.trim()) {
      onAddTask(newTaskTitle, newTaskDesc, newTaskStatus);
      setNewTaskTitle("");
      setNewTaskDesc("");
      setNewTaskStatus("todo");
    }
  };

  const getTasksByStatus = (status) => {
    return tasks.filter((task) => task.status === status);
  };

  return (
    <div className="w-full">
      <div className="card p-6 mb-8">
        <div className="flex items-center gap-3 mb-4">
          <div className="flex h-9 w-9 items-center justify-center rounded-md bg-blue-600 text-white">
            <FaPlus size={16} />
          </div>
          <h2 className="text-xl font-bold text-gray-900 m-0">
            Adicionar Nova Tarefa
          </h2>
        </div>

        <form onSubmit={handleAddTask} className="space-y-4">
          <div className="flex flex-col gap-4 lg:flex-row lg:items-end">
            <input
              type="text"
              placeholder="Título da tarefa..."
              value={newTaskTitle}
              onChange={(e) => setNewTaskTitle(e.target.value)}
              className="input-field flex-1"
              required
            />
            <textarea
              placeholder="Descrição (opcional)"
              value={newTaskDesc}
              onChange={(e) => setNewTaskDesc(e.target.value)}
              rows="2"
              className="input-field flex-1 resize-none"
            />
            <select
              value={newTaskStatus}
              onChange={(e) => setNewTaskStatus(e.target.value)}
              className="input-field min-w-[170px]"
            >
              {statuses.map((status) => (
                <option key={status} value={status}>
                  {statusLabels[status]}
                </option>
              ))}
            </select>
            <button type="submit" className="btn-primary whitespace-nowrap">
              Adicionar Tarefa
            </button>
          </div>
        </form>
      </div>

      <div className="card">
        <div className="grid grid-cols-1 md:grid-cols-3 divide-y md:divide-y-0 md:divide-x divide-gray-200 min-h-[600px]">
          {statuses.map((status) => {
            const Icon = statusIcons[status];

            return (
              <Column
                key={status}
                status={status}
                label={statusLabels[status]}
                icon={Icon}
                tasks={getTasksByStatus(status)}
                onUpdateTask={onUpdateTask}
                onDeleteTask={onDeleteTask}
                onMoveTask={onMoveTask}
              />
            );
          })}
        </div>
      </div>
    </div>
  );
}

export default KanbanBoard;
