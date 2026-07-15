export const fetchEvents = async () => {
  const response = await fetch('http://localhost:8080/api/v1/events');
  if (!response.ok) throw new Error('Failed to fetch events');
  return await response.json();
};

export const saveEvent = async (eventData) => {
  const response = await fetch('http://localhost:8080/api/v1/events', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(eventData),
  });
  return response.ok;
};

export const updateEventStatus = async (eventId, status) => {
  const response = await fetch(`http://localhost:8080/api/v1/events/${eventId}/status`, {
    method: 'PATCH',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ status }),
  });
  return response.ok;
};

export const deleteEvent = async (eventId) => {
  const response = await fetch(`http://localhost:8080/api/v1/events/${eventId}`, {
    method: 'DELETE',
  });
  return response.ok;
};
