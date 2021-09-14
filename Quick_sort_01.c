#include<stdio.h>
void quicksort(int number[25],int first,int last)
	{
		int i,j,pivot,temp;
		if(first<last)
			{
				pivot = first;
				i = first;
				j = last;
				while(i<j)
					{
						while(number[i] <= number[pivot] && i<last)
							i++;
						while(number[j] > number[pivot])
							j--;
						if(i<j)
							{
								temp = number[i];
								number[i] = number[j];
								number[j] = temp;
							}
					}
				temp = number[pivot];
				number[pivot] = number[j];
				number[j] = temp;
				quicksort(number,first,j-1);
				quicksort(number, j+1, last);
			}
	}
int main()
{
	int i,count;
	count = 5;	// number of elements in the array
	int number[5] = {23,54,15,96,57};
	quicksort(number,0,count-1);

	printf("\n Sorted elements:");
	for (i-0;i<count;i++)
		printf("%d",number[i]);
	return 0;
}


