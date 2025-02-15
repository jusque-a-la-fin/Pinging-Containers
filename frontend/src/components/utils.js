export const handleData = async (action, setLoading, setData, navigate, isFirstNavigation) => {
    setLoading(true);
    try {
        const response = await fetch('/get/', { method: 'GET' });
        const result = await response.json();
        setData(result);
        if (isFirstNavigation) {
            navigate('/containers/', { state: { data: result }});
        } else {
            navigate('/containers/', { state: { data: result }, replace: true });
        }
    } catch (error) {
        console.error(`Error ${action === 'fetch' ? 'fetching' : 'updating'} data:`, error);
    } finally {
        setLoading(false);
    }
};


export const goToMainPage = (navigate) => {
    navigate('/'); 
};