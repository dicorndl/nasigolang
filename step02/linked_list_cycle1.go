/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func hasCycle(head *ListNode) bool {
	if head == nil {
			return false
	}
	
	rabbit := head
	turtle := head
	for {
			if rabbit.Next == nil || rabbit.Next.Next == nil {
					return false
			}
			if turtle.Next == nil {
					return false
			}
			
			rabbit = rabbit.Next.Next
			turtle = turtle.Next
			
			if rabbit == turtle {
					return true
			}   
	} 
}