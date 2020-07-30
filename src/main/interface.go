package main

type dhtNode interface {
	/* "Run" is called after calling "NewNode". */
	Run()

	/* "Create" and "Join" are called after calling "Run". */
	/* For a dhtNode, either "Create" or "Join" will be called, but not both. */
	Create()               /* Create a new network. */
	Join(addr string) bool /* Join an existing network. Return "true" if join succeeded and "false" if not. */

	/* Quit from the network it is currently in.*/
	/* "Quit" will not be called before "Create" or "Join". */
	/* For a dhtNode, "Quit" may be called for many times. */
	/* For a quited node, call "Quit" again should have no effect. */
	Quit()

	/* Chord offers a way of "normal" quitting. */
	/* For "force quit", the node quit the network without informing other nodes. */
	/* "ForceQuit" will be checked by TA manually. */
	ForceQuit()

	/* Check whether the node represented by the IP address is in the network. */
	Ping(addr string) bool

	/* Put a key-value pair into the network (if KEY is already in the network, cover it), or
	 * get a key-value pair from the network, or
	 * remove a key-value pair from the network.
	 */
	Put(key string, value string) bool /* Return "true" if success, "false" otherwise. */
	Get(key string) (bool, string)     /* Return "true" and the value if success, "false" otherwise. */
	Delete(key string) bool            /* Remove the key-value pair represented by KEY from the network. */
	/* Return "true" if remove successfully, "false" otherwise. */
}
