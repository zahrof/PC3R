
public class Counter
{
	int myCount;

	public Counter()
	{
		myCount = 0;
	}

	public Counter(int init)
	{
		myCount = init;
	}

	public int getValue()
	{
		return myCount;
	}

	public void decrement()
	{
		synchronized(this) {
		myCount--;
		}
	}

}