using System.Collections;
using System.Collections.Generic;
using UnityEngine;
namespace DesignPattern2
{
    public class Singleton : MonoBehaviour
    {
        private static Singleton instance = new Singleton();
        public static Singleton Instance
        {
            get
            {
                if (instance == null)
                {
                    instance = new GameObject(nameof(Singleton)).AddComponent<Singleton>();
                }
                return instance;
            }
        }
    }
}
