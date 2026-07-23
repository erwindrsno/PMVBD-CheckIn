export const fetchEvents = async () => {
  const response = await fetch('http://localhost:8080/api/v1/events', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${sessionStorage.getItem('token')}`
    }
  });
  if (response.status === 401) {
    sessionStorage.removeItem('token');
    window.location.href = '/login'; // Or use router.push('/login')
    return;
  }
  if (!response.ok) throw new Error('Failed to fetch events');
  return await response.json();
};

export const saveEvent = async (eventData) => {
  const response = await fetch('http://localhost:8080/api/v1/events', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${sessionStorage.getItem('token')}`
    },
    body: JSON.stringify(eventData),
  });
  if (response.status === 401) {
    sessionStorage.removeItem('token');
    window.location.href = '/login'; // Or use router.push('/login')
    return;
  }
  return response.ok;
};

export const updateEventStatus = async (eventId, status) => {
  const response = await fetch(`http://localhost:8080/api/v1/events/${eventId}/status`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${sessionStorage.getItem('token')}`
    },
    body: JSON.stringify({ status }),
  });
  if (response.status === 401) {
    sessionStorage.removeItem('token');
    window.location.href = '/login'; // Or use router.push('/login')
    return;
  }
  return response.ok;
};

export const deleteEvent = async (eventId) => {
  const response = await fetch(`http://localhost:8080/api/v1/events/${eventId}`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${sessionStorage.getItem('token')}`
    },
  });
  if (response.status === 401) {
    sessionStorage.removeItem('token');
    window.location.href = '/login'; // Or use router.push('/login')
    return;
  }
  return response.ok;
};
