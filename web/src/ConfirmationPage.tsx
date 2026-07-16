import { useNavigate, useParams } from "react-router-dom";

export const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080/v1";

export function ConfirmationPage() {
  const { token } = useParams();
  const navigate = useNavigate();

  const handleConfirm = async () => {
    const response = await fetch(`${API_URL}/users/activate/${token}`, {
      method: "PUT",
    });

    if (response.ok) {
      navigate("/");
    } else {
      alert("Failed to confirm token");
    }
  };

  return (
    <div>
      <h1>Confirmation</h1>
      <button onClick={handleConfirm}>Click to confirm</button>
    </div>
  );
}
