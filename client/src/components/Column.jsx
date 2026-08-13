import { useState } from "react";
import { FaSave, FaTimes } from "react-icons/fa";
import TaskCard from "./TaskCard";

function Column({
  status,
  label,
  icon: Icon,
  tasks,
  onUpdateTask,
  onDeleteTask,
  onMoveTask,
}) {
  const [editingId, setEditingId] = useState(null);
  const [editTitle, setEditTitle] = useState("");
  const [editDesc, setEditDesc] = useState("");

  const handleEditStart = (task) => {
    setEditingId(task.id);
    setEditTitle(task.title);
    setEditDesc(task.description || "");
  };

  const handleEditSave = (taskId) => {
    onUpdateTask(taskId, {
      title: editTitle,
      description: editDesc,
      status: status,
    });
    setEditingId(null);
  };

  const handleEditCancel = () => {
    setEditingId(null);
  };

  const moveTask = (taskId, direction) => {
    const statuses = ["todo", "doing", "done"];
    const currentIndex = statuses.indexOf(status);

    if (direction === "next" && currentIndex < statuses.length - 1) {
      onMoveTask(taskId, statuses[currentIndex + 1]);
    } else if (direction === "prev" && currentIndex > 0) {
      onMoveTask(taskId, statuses[currentIndex - 1]);
    }
  };

  return (
    <div className="flex flex-col h-full">
      <div className="bg-gray-50 border-b border-gray-200 px-4 py-3 flex justify-between items-center flex-shrink-0">
        <div className="flex items-center gap-2">
          <div className="flex h-5 w-5 items-center justify-center text-blue-600">
            {Icon && <Icon size={13} />}
          </div>
          <h3 className="text-sm font-bold text-gray-900 m-0">{label}</h3>
        </div>
        <span className="bg-gray-200 px-2 py-0.5 rounded text-xs font-semibold text-gray-600">
          {tasks.length}
        </span>
      </div>

      <div className="flex-1 overflow-y-auto p-3 flex flex-col">
        {tasks.length === 0 ? (
          <div className="flex flex-col items-center justify-center flex-1 text-center text-gray-400 py-6">
            <p className="text-sm font-medium mb-1">Vazio</p>
            <small className="text-xs">Nenhuma tarefa aqui</small>
          </div>
        ) : (
          <div className="flex flex-col gap-2">
            {tasks.map((task) =>
              editingId === task.id ? (
                <div
                  key={task.id}
                  className="bg-blue-50 p-3 rounded border-2 border-blue-300 flex flex-col gap-2"
                >
                  <input
                    type="text"
                    value={editTitle}
                    onChange={(e) => setEditTitle(e.target.value)}
                    placeholder="Título"
                    className="input-field"
                  />
                  <textarea
                    value={editDesc}
                    onChange={(e) => setEditDesc(e.target.value)}
                    placeholder="Descrição"
                    rows="2"
                    className="input-field resize-none"
                  />
                  <div className="flex gap-2">
                    <button
                      className="flex-1 px-2 py-1.5 bg-green-600 text-white rounded font-semibold hover:bg-green-700 transition-colors duration-200 flex items-center justify-center gap-2 text-xs"
                      onClick={() => handleEditSave(task.id)}
                    >
                      <FaSave size={12} />
                      Salvar
                    </button>
                    <button
                      className="flex-1 px-2 py-1.5 bg-gray-300 text-gray-700 rounded font-semibold hover:bg-gray-400 transition-colors duration-200 flex items-center justify-center gap-2 text-xs"
                      onClick={handleEditCancel}
                    >
                      <FaTimes size={12} />
                      Cancelar
                    </button>
                  </div>
                </div>
              ) : (
                <TaskCard
                  key={task.id}
                  task={task}
                  onEdit={() => handleEditStart(task)}
                  onDelete={() => onDeleteTask(task.id)}
                  onMovePrev={() => moveTask(task.id, "prev")}
                  onMoveNext={() => moveTask(task.id, "next")}
                  canMovePrev={["todo", "doing", "done"].indexOf(status) > 0}
                  canMoveNext={["todo", "doing", "done"].indexOf(status) < 2}
                />
              ),
            )}
          </div>
        )}
      </div>
    </div>
  );
}

export default Column;
