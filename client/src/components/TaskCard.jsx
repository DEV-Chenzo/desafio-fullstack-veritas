import { FaEdit, FaTrash, FaChevronLeft, FaChevronRight } from "react-icons/fa";

function TaskCard({
  task,
  onEdit,
  onDelete,
  onMovePrev,
  onMoveNext,
  canMovePrev,
  canMoveNext,
}) {
  const formatDate = (date) => {
    return new Date(date).toLocaleDateString("pt-BR", {
      month: "2-digit",
      day: "2-digit",
      hour: "2-digit",
      minute: "2-digit",
    });
  };

  return (
    <div className="bg-white border border-gray-200 rounded p-3 hover:border-gray-300 transition-colors duration-200">
      <div className="flex justify-between items-start gap-2 mb-2">
        <h4 className="text-sm font-semibold text-gray-900 flex-1 break-words m-0">
          {task.title}
        </h4>
        <button
          className="flex h-7 w-7 items-center justify-center text-gray-400 hover:text-red-600 transition-colors duration-200 flex-shrink-0"
          onClick={onDelete}
          title="Deletar"
        >
          <FaTrash size={12} />
        </button>
      </div>

      {task.description && (
        <p className="m-0 mb-2 text-xs text-gray-600 leading-relaxed break-words">
          {task.description}
        </p>
      )}

      <div className="mb-2 border-t border-gray-100 pt-1.5">
        <small className="text-xs text-gray-400">
          {formatDate(task.createdAt)}
        </small>
      </div>

      <div className="flex gap-1">
        <button
          className="flex items-center justify-center gap-1 flex-1 px-2 py-1 border border-blue-300 bg-blue-50 text-blue-700 rounded text-xs font-medium hover:bg-blue-100 transition-colors duration-200 whitespace-nowrap"
          onClick={onEdit}
          title="Editar tarefa"
        >
          <FaEdit size={10} />
          <span>Editar</span>
        </button>

        {canMovePrev && (
          <button
            className="flex items-center justify-center px-2 py-1 border border-gray-300 bg-gray-50 text-gray-700 rounded text-xs font-medium hover:bg-gray-100 transition-colors duration-200"
            onClick={onMovePrev}
            title="Mover para coluna anterior"
          >
            <FaChevronLeft size={10} />
          </button>
        )}

        {canMoveNext && (
          <button
            className="flex items-center justify-center px-2 py-1 border border-gray-300 bg-gray-50 text-gray-700 rounded text-xs font-medium hover:bg-gray-100 transition-colors duration-200"
            onClick={onMoveNext}
            title="Mover para próxima coluna"
          >
            <FaChevronRight size={10} />
          </button>
        )}
      </div>
    </div>
  );
}

export default TaskCard;
